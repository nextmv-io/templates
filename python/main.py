"""
Template for working with Python.
"""

import argparse
import json
import sys
import time
from typing import Any

from nextmv import __about__


def main() -> None:
    """Entry point for the template."""

    parser = argparse.ArgumentParser(description="Template for working with Python.")
    parser.add_argument(
        "-input",
        default="",
        help="Path to input file. Default is stdin.",
    )
    parser.add_argument(
        "-output",
        default="",
        help="Path to output file. Default is stdout.",
    )
    parser.add_argument(
        "-recipient",
        default="world",
        help="Who do you want to greet?",
        type=str,
    )
    args = parser.parse_args()

    # Read input data, solve the problem and write the solution.
    input_data = read_input(args.input)
    log(f'{input_data["greeting"]}, {args.recipient}!')

    solution = solve(input_data, args.recipient)
    write_output(args.output, solution)


def solve(input_data: dict[str, Any], recipient: str) -> dict[str, Any]:
    """Solve your model here."""

    start_time = time.time()

    greetingMessage = f'{input_data["greeting"]}, {recipient}!'

    end_time = time.time()
    duration = end_time - start_time

    # Creates the statistics.
    statistics = {
        "result": {
            "custom": {
                "greeting_message": greetingMessage,
                "nextmv_version": __about__.__version__,
            },
            "duration": duration,
            "value": len(greetingMessage),
        },
        "run": {
            "duration": duration,
        },
        "schema": "v1",
    }

    return {
        "solutions": [{"greeting_message": greetingMessage}],
        "statistics": statistics,
    }


def log(message: str) -> None:
    """Logs a message. We need to use stderr since stdout is used for the
    solution."""

    print(message, file=sys.stderr)


def read_input(input_path) -> dict[str, Any]:
    """Reads the input from stdin or a given input file."""

    input_file = {}
    if input_path:
        with open(input_path, encoding="utf-8") as file:
            input_file = json.load(file)
    else:
        input_file = json.load(sys.stdin)

    return input_file


def write_output(output_path, output) -> None:
    """Writes the output to stdout or a given output file."""

    content = json.dumps(output, indent=2)
    if output_path:
        with open(output_path, "w", encoding="utf-8") as file:
            file.write(content + "\n")
    else:
        print(content)


if __name__ == "__main__":
    main()
