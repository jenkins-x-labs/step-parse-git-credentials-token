# step-get-git-credentials-token

A Jenkins X Labs repository that looks to work with Git credentials in a pipeline.

The main aim we are trying here is to parse a local `~/.git-credentials` file rather than
accessing Kubernetes secrets directly which requires increased roles and means the pipeline 
can access all secrets in the namespace.  Using this step means we use a locally generated
file to read secrets. 

This could probably have been done using sed but for some go is easier :)