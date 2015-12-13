Browser-to-browser chat example
===============================

This example uses https://github.com/nsf/bin2go to embed assets. To
make sure it's installed, run

    go get github.com/nsf/bin2go
    
Other neccessary packages

    go get -v github.com/gorilla/websocket
    go get -v github.com/tv42/birpc
    go get -v github.com/tv42/topic

and put the resulting bin2go in `PATH`.

To run the example:

	./run

or, to select another port:

    ./run -port=8888

Then open two browser windows to http://localhost:8000/ , change your
name from *J. Doe* to something more distinct, and type messages. See
how they are transmitted to the other browser window.

See the browser's Javascript console for a debug log of incoming and
outgoing messages.
