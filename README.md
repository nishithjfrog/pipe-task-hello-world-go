# JFrog Pipelines Task Hello World

This repo contains a basic Pipelines Task and Pipelines definitions to help test and publish it. 
Developers can use this to learn about Pipelines Tasks and as a reference to create their own Pipelines tasks.

The task available in this repo was developed using the Pipelines Tasks GO SDK. 
It shows how to use the following features available to Pipelines Tasks:

- Read task input arguments.
- Log info and error messages.
- Create new file at current step working directory.
- Save and restoring information to task state.
- Export new environment variable.
- Set task output values.
- Using task post hooks.

Pipelines definitions are also available to help developers with the following activities:

- Build and run Pipelines Task from source to validate changes.
- Build and publish Pipelines Task to Artifactory.
- Run task from Artifactory to validate published task.

### Pipelines Tasks in Action!

To see this Pipelines Task in action, do the following:

- Fork this repo.
- Change [values.yml](.jfrog-pipelines/values.yml) contents to reflect your settings and source code location.
- Add a Pipelines Source pointing to your forked repo.

After Pipelines Sync succeeds you should see and trigger the following Pipelines:

- **hello_world_task_test**: This Pipeline can be used to run this Pipelines Task from source on a Bash step and check
if the task is producing the expected behavior. This pipeline will trigger automatically every time you push changes to
the source code.
- **hello_world_task_publish**: This Pipeline can be used to build and publish the task to Artifactory. 
After publishing is done it will also run the recently published task to make sure the published package works as expected
and can be used by others.

## Important Files

- [task.yml](task.yml): Pipelines Task descriptor file. Contains the details 
about the task, including name, description, inputs, outputs and commands to be executed.
- [main.go](main.go): The main entrypoint where the task logic starts.
- [Makefile](Makefile): This defines how the binaries are created. They are specifically formatted to be usable on all platforms that are supported by Pipelines. See the task.yaml for how to contextually choose a binary to execute.
- [test-task-pipeline.yaml](.jfrog-pipelines/test-task-pipeline.yaml): Pipelines definition declaring
the _hello_world_go_task_test_ pipeline.
- [publish-task-pipeline.yaml](.jfrog-pipelines/publish-task-pipeline.yaml): Pipelines definition declaring
the _hello_world_go_task_publish_ pipeline.

## Additional References

- [Using the shared development environment](https://docs.google.com/document/d/1c3p49HSGAUzFO0pfV4fArJ_uM7BdE7-b6Dtp8lyLci4/edit?usp=sharing)
