# Nextmv Java template

This template helps you writing your own optimization model in Java.
It contains several files to easily deploy the model to the Nextmv
platform.

## Usage

Java templates in Nextmv require a `main.jar` as an entry point. Running the
following command will generate a `main.jar` in the root direcotry of the
project.

```bash
mvn package
```

After that you can run the `main.jar` file with the following command:

```bash
java -jar main.jar --input input.json
```

You can also push the `main.jar` file to the Nextmv Cloud and run it remotely.
Take a look at the documentation on how to
[deploy](https://www.nextmv.io/docs/platform/deploy-app/custom-apps) and
[run](https://www.nextmv.io/docs/platform/run-app-remotely/nextmv-cli) an app in
the Nextmv Cloud.

## Mirror running on Nextmv Cloud locally

Pre-requisites: Docker needs to be installed.

To run the application locally in the same docker image as the one used on the
Nextmv Cloud, you can use the following command:

```bash
mvn package && cat input.json | docker run -i --rm \
-v $(pwd):/app ghcr.io/nextmv-io/runtime/java:latest \
java -jar /app/main.jar
```

You can also debug the application by running it in a Dev Container. This
workspace recommends to install the Dev Container extension for VSCode. If you
have the extension installed, you can open the workspace in a container by using
the command `Dev Containers: Reopen in Container`.

## Next steps

* Open `main.jar` and start writing the model in Java.
* Further documentation, guides, and API references about custom modeling and
  deployment can also be found on our [blog](https://www.nextmv.io/blog) and on
  our [documentation site](https://docs.nextmv.io).
* Need more assistance? Send us an [email](mailto:support@nextmv.io)!
