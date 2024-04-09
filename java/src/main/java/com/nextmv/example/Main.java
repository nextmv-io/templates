package com.nextmv.example;

import java.util.ArrayList;
import java.util.List;

public final class Main {

  public static void main(String[] args) {
    // Parse arguments. Exit on error.
    Options options = Options.fromArguments(args);

    // Load input. Exit on error.
    Input input = Input.fromString(options.getInputPath());

    long startTime = System.currentTimeMillis();

    // Build the output message.
    String greetingMessage = input.getGreeting() + ", " + options.getRecipient() + "!";

    long endTime = System.currentTimeMillis();
  
    Output output = new Output(
        greetingMessage,
        (endTime - startTime) / 1000.0,
        greetingMessage.length()
    );
    
    // Write output. Exit on error.
    Output.write(options.getOutputPath(), output);
  }
}
