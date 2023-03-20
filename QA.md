# Q&As


## For the persistent storage, should we use a database image, or should we connect to an external database. I would think we need to use an external one, this way we are actually persisting the data between deployments, but I'm having networking problems on the cluster.

Given our minimal Kubernetes configuration, the instruction "persist data" in Option 6 should not be interpreted as creating a persistent volume claim in Kubernetes. Instead, you are supposed to use a database image and set up a database deployment in your solution, even though the data is not actually persisted. Please avoid using an external database since there may be some connection and consistency issues when evaluating your solution.
