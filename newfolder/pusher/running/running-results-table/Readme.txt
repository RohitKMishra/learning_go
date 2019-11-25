Building the backend service
We are going to write our backend service using the Go language, using the Gin library to power our HTTP service.

Our service is going to offer two endpoints:

GET /results - this returns the current list of results.
POST /results - this creates a new result to add to the list.
To start with, we need to create an area to work with. Create a new directory under your GOPATH in which to work:


We can then initialise our work area for this project. This is done using the dep tool:

    $ dep init
Doing this will create the *Gopkg.toml and Gopkg.lock files used to track our dependencies, and the vendor *directory which is used to store vendored dependencies.

The next thing to do is to create our data store. We are going to do this entirely in memory for this article, but in reality you would use a real database, for example PostgreSQL or MongoDB.

Create a new directory called internal/db under our work area, and create a db.go file in here as follows:

Here we are creating a new type called Record that represents the data that we store, and a new struct called Database that represents the actual database we are using. We then create some methods on the Database type to add a record and to get the list of all records.

Next we can create our web server. For this we are going to create a new directory called internal/webapp under our work area, and a new file called webapp.go in this directory as follows:

This creates a function called StartServer that will create and run our web server, defining two routes on it to do the processing that we need.

We are also importing some packages that aren’t currently available - github.com/gin-gonic/gin and github.com/gin-contrib/cors. The first of these is the Gin web server itself, and the second is the contrib library to enable CORS, so that our webapp can access the backend server.

We can now use dep to ensure that this is available for us, by executing dep ensure from our top level. This will download the necessary packages and put them into our vendor directory ready to be used:

Finally, we create a main program that actually makes use of this all. For this, in the top of the work area we create a file called running-results-table.go as follows:

This makes use of our db and webapp modules that we’ve just written, and starts everything up correctly.

We can now run our application by executing go run running-results-table.go:

Alternatively, we can build an executable using go build running-results-table.go. This executable can then be distributed however we need to do so - for example, copying it into a Docker container or directly onto our production VMs.


Sending live updates when data changes
At this point, we can correctly create new records and retrieve all of the records that have been created. However, there is no support for live updates at this point - the client would need to keep re-requesting the data to see if anything changes.

As a better solution to this, we are going to use Pusher Channels to automatically emit events whenever a new record is created, so that all listening clients can automatically update themselves without needing to poll the server. Additionally, we are going to use Go channels to isolate the sending of Pusher events from the actual HTTP request - allowing our server to respond to the client faster, whilst still sending the event a fraction of a second later.

Create a new directory called internal/notifier under our work area, and in this create a file called notifier.go as follows:

There is quite a lot going on here, so lets work through it.

The first thing we do is define a new type called Notifier. This is our interface that we expose to the rest of the code through which we can notify clients of new results.

Next, we define a non-exported function called notifier that is given a reference to the database and a Go channel. This function will create our Pusher client, and then start an infinite loop of reading from the channel (which blocks until a new message comes in), retrieving the latest list of results from the database and sending them off to Pusher. We deliberately get the latest list ourselves here in case there was some delay in processing the message - this way we’re guaranteed not to miss anything.

We then create a new method called New that will return a new Notifier. Importantly in here we also start a new go-routine that runs our notifier function, which essentially means that there is a new thread of execution running that function.

Finally we have a Notify method on our Notifier that does nothing more than push a new value down our Go channel.

The end result of this is that, whenever someone calls Notifier.Notify(), we will trigger our go-routine - on a separate thread - to retrieve the current results from the database and send them to Pusher.

We now need to use dep to again ensure that this is available for us, by executing dep ensure from our top level.