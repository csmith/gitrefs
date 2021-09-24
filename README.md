# gitrefs

Provides a simple way to list references from a remote git repository over HTTP, without needing a full git client or a
checkout of the remote repository.

Example:

```go
package main

import (
	"fmt"

	"github.com/csmith/gitrefs"
)

func main() {
	refs, err := gitrefs.Fetch("https://github.com/csmith/gitrefs")
	if err != nil {
		panic(err)
	}

	for r := range refs {
		fmt.Printf("Ref %s at commit %s\n", r, refs[r])
	}
}
```

A utility method to retrieve only the latest semver tag is included:

```go
package main

import (
	"fmt"

	"github.com/csmith/gitrefs"
)

func main() {
	tag, hash, err := gitrefs.LatestTag("https://github.com/csmith/gitrefs")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Latest tag is %s at commit %s\n", tag, hash)
}
```

Protocol docs:
- https://www.git-scm.com/docs/http-protocol
- https://git-scm.com/docs/pack-protocol
