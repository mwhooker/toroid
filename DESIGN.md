# Toroid design doc


CI should be simple. The idea is we want git pushes to trigger user defined jobs which run as quickly as possible. What we really want is something like travisci, where a .travis file defines what's run when a repo is pushed. We also want Travis' workspace isolation so that the host OS isn't where library dependencies are installed.

Looking at where we want to be clearly shows cgroups being used for containerization during the build process. The difference between what we want and what Travis provides is that in certain cases we want to export our build environment as the build artifact, ie an ami.

Another slight difference is we want a build pipeline. Whereas Travis is good for running unit tests, we want a successful test run to cascade into a deploy to an environment.

Putting it all together we have what I consider to be an incredibly simple system.

Here's a workflow.

### Library build
A library produces an artifact and nothing else.

1. A push to the library's master branch
2. Github puts a message on sqs
3. CI service picks up the message.
4. Checks out library.
5. Builds it.

**build**

Building will happen inside of a container created by packer.
If a .packer file exists, it will be used. If not, one will be created.

Buidling will, for example, run tests and upload the library to rubygems

### Application.
An application build implies a pipeline.

Same as above except the artifact is an ami. At the end it puts a new message on the queue saying to run the build for the next environment. The environment is arbitrary. Let's say its stage. The build would for example involve creating a new cloud formation stack with the previously created AMI for the stage environment. Upon a successful deploy

**build**
The build scripts can be implemented in any way, but the interface must be consistent. There must be an entry point for each supported environment (could be parametrized). There must also be syntax for populating the build queue.



This is a nice basic system. A secondary goal would be adding some small amount if state and a web ui. That way we could inspect builds like in Travis or pause certain environments.

The work for this is underway already. Since I believe v1 will have a small footprint, I believe building it first and iterating is the right move.
