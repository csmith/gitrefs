package gitrefs

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const (
	refsPath    = "info/refs?service=git-upload-pack"
	contentType = "application/x-git-upload-pack-advertisement"
	header      = "# service=git-upload-pack"
	tagPrefix   = "refs/tags/"
)

type opts struct {
	ctx      context.Context
	client   *http.Client
	tagsOnly bool
}

type Option func(*opts)

func HttpClient(client *http.Client) Option {
	return func(o *opts) {
		o.client = client
	}
}

func TagsOnly() Option {
	return func(o *opts) {
		o.tagsOnly = true
	}
}

func WithContext(ctx context.Context) Option {
	return func(o *opts) {
		o.ctx = ctx
	}
}

// Fetch retrieves a list of all refs from the remote repository at the given url.
func Fetch(url string, options ...Option) (map[string]string, error) {
	o := &opts{
		ctx:      context.Background(),
		client:   http.DefaultClient,
		tagsOnly: false,
	}

	for i := range options {
		options[i](o)
	}

	req, err := http.NewRequestWithContext(o.ctx, http.MethodGet, fmt.Sprintf("%s/%s", url, refsPath), nil)
	if err != nil {
		return nil, err
	}

	res, err := o.client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected response code: %d", res.StatusCode)
	}

	if ct := res.Header.Get("Content-Type"); ct != contentType {
		return nil, fmt.Errorf("unexpected content type: %s", ct)
	}

	defer res.Body.Close()
	refs := make(map[string]string)
	body := false
	for {
		line, flush, err := readPktLine(res.Body)
		if err != nil {
			return nil, fmt.Errorf("unable to decode pkt-line: %v", err)
		}

		if flush {
			if body {
				return refs, nil
			} else {
				body = true
			}
		} else if !body {
			if strings.TrimSpace(string(line)) != header {
				return nil, fmt.Errorf("invalid header received: %s", line)
			}
		} else {
			parts := strings.SplitN(string(line), " ", 2)
			hash := parts[0]
			ref := strings.TrimSpace(strings.SplitN(parts[1], "\000", 2)[0])
			if strings.HasSuffix(ref, "^{}") {
				continue
			}
			if !o.tagsOnly || strings.HasPrefix(ref, tagPrefix) {
				refs[ref] = hash
			}
		}
	}
}

func readPktLine(reader io.Reader) ([]byte, bool, error) {
	lengthHex := make([]byte, 4)
	if _, err := io.ReadFull(reader, lengthHex); err != nil {
		return nil, false, err
	}

	length, err := strconv.ParseUint(string(lengthHex), 16, 64)
	if err != nil {
		return nil, false, err
	}

	if length == 0 {
		// "Flush" packet
		return nil, true, err
	}

	line := make([]byte, length-4)
	if _, err := io.ReadFull(reader, line); err != nil {
		return nil, false, err
	}

	return line, false, nil
}
