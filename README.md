# NAUTICA
NAUTICA is a web service for dive sequence calculation based on US Navy No-Decompression Dive Table revision 7.


### Tech

Nautica  uses a number of open source projects to work properly:

* [golang] - The Go programming language .
* [gorilla-mux] - A powerful URL router and dispatcher for golang.

And of course NAUTICA itself is open source with a [public repository][dill]
 on GitHub.

## Installation

NAUTICA requires [Gorilla-mux](https://github.com/gorilla/mux) to run.

### Install the Go tools
Download the archive and extract it into /usr/local, creating a Go tree in /usr/local/go.

    wget https://redirector.gvt1.com/edgedl/go/go1.9.2.linux-amd64.tar.gz
    sudo tar -C /usr/local -xzf go1.9.2.linux-amd64.tar.gz

Add /usr/local/go/bin to the PATH environment variable. You can do this by adding this line to your /etc/profile (for a system-wide installation) or $HOME/.profile:

    export PATH=$PATH:/usr/local/go/bin

### Install Git, if needed
To perform the next step you must have Git installed. (Check that you have a git command before proceeding.)

If you do not have a working Git installation, follow the instructions on the [Git downloads](http://git-scm.com/downloads) page.

### Clone gorilla/mux
    go get https://github.com/gorilla/mux.git
    
### Clone NAUTICA
    go get https://github.com/18215036/nautica.git
    
## How to Run    
    cd $HOME/go/src/github.com/18215036/nautica/
    go run main.go

[//]: # 
   [dill]: <https://github.com/18215036/nautica>
   [git-repo-url]: <https://github.com/18215036/nautica.git>
