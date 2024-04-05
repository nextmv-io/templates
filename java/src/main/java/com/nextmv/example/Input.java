package com.nextmv.example;

import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.util.stream.Collectors;

import com.google.gson.Gson;

public class Input {
  private final String greeting;

  public Input(String greeting) {
    this.greeting = greeting;
  }

  public String getGreeting() {
    return this.greeting;
  }
  
  public static Input fromString(String path) {
    Gson gson = new Gson();
    // Read stdin if no path is provided.
    if (path.isEmpty()) {
      try (BufferedReader reader = new BufferedReader(new InputStreamReader(System.in))) {
        return gson.fromJson(
          reader.lines().collect(Collectors.joining("\n")), Input.class
        );
      } catch (java.io.IOException e) {
        System.err.println("Error reading stdin: " + e.getMessage());
        System.exit(1);
        return null;
      }
    }
    // Read the path otherwise.
    try {
      return gson.fromJson(
        java.nio.file.Files.readString(java.nio.file.Paths.get(path)),
        Input.class
      );
    } catch (java.io.IOException e) {
      System.err.println("Error reading '" + path + "': " + e.getMessage());
      System.exit(1);
    }
    return null;
  }
}
