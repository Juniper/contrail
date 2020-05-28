"""
Tool to generate requirements.txt from requirements.in file
using pip-compile https://pypi.org/project/pip-tools/

Tool accepts requirements.in with only top/first level dependencies
and generates requirements.txt with recursive pinned down
dependencies or top/first level dependencies

Generates new requirements.txt, only if requirements.in
is modified.

To force generate requirements.txt irrespective of change
in requirements.in use --force|-f commandline argument

Returns non zero return code until the generated requirements.txt
is committed to local git branch
"""

import os
import sys
import argparse
import subprocess
import requirements

DEFAULT_REQUIREMENTS_INPUT_FILE = "requirements.in"
DEFAULT_REQUIREMENTS_OUTPUT_FILE = "requirements.txt"


def parse_args():
    parser = argparse.ArgumentParser(
            description='Generate requirments.txt from requirments.in')
    parser.add_argument('-i', '--inreq', default='requirements.in',
                        help='First level requirements.in input file')
    parser.add_argument('-o', '--outreq', default='requirements.txt',
                        help='Generated requirements.txt output file')
    parser.add_argument('-f', '--force', default=False,
                        help='Force generate requirements.txt')

    return parser.parse_args()


def print_commit_error(outreq):
    """Print message informing users to commit the modified
    requirements.txt
    """
    print('*'*80)
    print("%s is changed, please commit it to git repository" % outreq)
    print('*'*80)
    sys.exit(1)


def is_committed(outreq):
    """Check whether the modified requirements.txt
    is committed to git repository
    """
    git_status_cmd = ["git status", "--porcelain", "--", outreq]
    if subprocess.check_output(' '.join(git_status_cmd), shell=True):
        git_diff_cmd = ["git diff", "--", outreq]
        subprocess.call(' '.join(git_diff_cmd), shell=True)
        print_commit_error(outreq)


def generate(inreq, outreq):
    """Generate requirements.txt with given requirements.in
    using pip-compile tool
    """
    pipcompile_cmd = [
            'CUSTOM_COMPILE_COMMAND="python gen-pipreqs.py" pip-compile']
    pipcompile_args = [inreq, "--output-file " + outreq]
    pipcompile_cmd.extend(pipcompile_args)
    subprocess.check_call(' '.join(pipcompile_cmd), shell=True)
    print_commit_error(outreq)


def format_req_str(req):
    """Format dependency string from name and specs
    """
    specs = [comparator + version for comparator, version in req.specs]
    return req.name + ",".join(specs)


def main():
    args = parse_args()

    # force generate requirements.txt
    if args.force:
        print("Force generating %s" % args.outreq)
        generate(args.inreq, args.outreq)
        return

    # requirements.txt doesn't exist, generate it
    if not os.path.exists(args.outreq):
        print("Missing: %s generating it" % args.outreq)
        generate(args.inreq, args.outreq)
        return

    # Get list of existing first level dependencies from requirements.txt
    with open(args.outreq, 'r') as outfile:
        existing_reqs = set([])
        for req in requirements.parse(outfile):
            if args.inreq in req.line:  # First level dependency
                existing_reqs.add(format_req_str(req))

    # Get list of desired first level dependencies from requirements.in
    with open(args.inreq, 'r') as infile:
        desired_reqs = set([])
        for req in requirements.parse(infile):
            desired_reqs.add(format_req_str(req))

    # Check whether requirements.in contains new first level
    # dependency in requirements.in
    new_reqs = desired_reqs.difference(existing_reqs)
    if new_reqs:
        print("First level %s added/modified, generating %s" % (
              new_reqs, args.outreq))
        generate(args.inreq, args.outreq)

    # Check whether first level dependencies are removed
    # in requirements.in
    stale_reqs = existing_reqs.difference(desired_reqs)
    if stale_reqs:
        print("Frist level %s removed, generating %s" % (
              stale_reqs, args.outreq))
        generate(args.inreq, args.outreq)

    # requirements up to data, make sure it is committed to git repo
    is_committed(args.outreq)
    print("%s is up to date" % args.outreq)


if __name__ == "__main__":
    main()
