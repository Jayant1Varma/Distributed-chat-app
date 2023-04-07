# Q&As


## For the persistent storage, should we use a database image, or should we connect to an external database. I would think we need to use an external one, this way we are actually persisting the data between deployments, but I'm having networking problems on the cluster.

Given our minimal Kubernetes configuration, the instruction "persist data" in Option 6 should not be interpreted as creating a persistent volume claim in Kubernetes. Instead, you are supposed to use a database image and set up a database deployment in your solution, even though the data is not actually persisted. Please avoid using an external database since there may be some connection and consistency issues when evaluating your solution.

## The pacslab/chatapp example defines a simple chat message type with keys id, name, email, date, topic, and content. Are we allowed to define our own structure for the message (containing the required data) or do we have to extend that structure?
You may define your own data structure for messages. Please note certain information must be included in each message.

## Related to the previous question: I see in the instructions it says "Each message should at least contain the following information" (the keys mentioned above are listed here). Are we allowed to have different types of messages (possibly for transferring data such as upvotes or around chat rooms)?
Yes

## Can we create our own API/flow for the chat application?
Yes

## What parts will be auto-graded?
The automated judge is more like a CI/CD pipeline, which retrieves the code from your submission, builds your code (e.g., Docker images), pushes images, and deploys your solution (please refer to Evaluation in the project description). After your system is deployed on a Kubernetes cluster, the TAs will manually evaluate your solution (e.g., check whether features are working). Therefore, you should precisely follow the instructions (e.g., directory structure, port number, and credentials for the registry) in the project description.

## Do we need strict API requirements?
No

## Do we have to extend the example front-end, or can we completely rewrite it with our own design/logic?
You can do anything you see fit.

## As an aside, the project description mentioned "You can modify its source code and rewrite it in Go to implement the chat application for this project" -- so I'm thinking we can fully design this ourselves if we want, or optionally base it off the example.
Correct

## Do we have to name images as image-a, image-b, and image-c, or can we use other names?
The three images in the project repo are just placeholders. You can use any image name you see fit. Just make sure that you update the directory name accordingly. For example, if you have an image named `my-awesome-frontend`, you should place all the necessary source files for building the image in the `images/my-awesome-frontend` directory. When building your images, the automated judge will fetch the image name and tag from `chatapp.yaml` and locate the source files from the `images` directory.

## Can we have more than three images?
Yes
