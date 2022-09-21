#!/bin/bash

# Test config
CMD=$1
PROJECT_DIR="sample-project"

extract_sample_project() {
    tar -xzf $PROJECT_DIR.tar.gz
}

delete_sample_project() {
    rm -rf $PROJECT_DIR
}

validate_input() {
    # test files exist
    if [[ ! -f "$PROJECT_DIR/.hidden1.jpeg" ]]; then exit 1; fi
    if [[ ! -f "$PROJECT_DIR/1.jpeg" ]]; then exit 1; fi
    if [[ ! -f "$PROJECT_DIR/2.jpeg" ]]; then exit 1; fi
    if [[ ! -f "$PROJECT_DIR/3.mov" ]]; then exit 1; fi
    if [[ ! -f "$PROJECT_DIR/4.mov" ]]; then exit 1; fi
}

# assert that old files do not exist & new files with dated names exist
validate_renaming() {
    # .hidden1.jpeg
    if [[ ! -f "$PROJECT_DIR/.hidden1.jpeg" ]]; then exit 1; fi

    # 1.jpeg
    if [[ -f "$PROJECT_DIR/1.jpeg" ]]; then exit 1; fi
    if [[ ! -f "$PROJECT_DIR/2017_01_15_14_10_35.jpeg" ]]; then exit 1; fi
    # 2.jpeg
    if [[ -f "$PROJECT_DIR/2.jpeg" ]]; then exit 1; fi
    if [[ ! -f "$PROJECT_DIR/2021_05_23_08_05_12.jpeg" ]]; then exit 1; fi
    # 3.mov
    if [[ -f "$PROJECT_DIR/3.mov" ]]; then exit 1; fi
    if [[ ! -f "$PROJECT_DIR/2013_12_11_11_12_13.mov" ]]; then exit 1; fi
    # 4.mov
    if [[ -f "$PROJECT_DIR/4.mov" ]]; then exit 1; fi
    if [[ ! -f "$PROJECT_DIR/1999_08_07_19_25_47.mov" ]]; then exit 1; fi
}

run_test() {
    delete_sample_project
    extract_sample_project
    validate_input
    $CMD $PROJECT_DIR
    validate_renaming
    delete_sample_project
}

run_version_test() {
    output=$($CMD -version)
    if [[ ! "$output" =~ "version:" ]]; then exit 1; fi
}

run_test

run_version_test

exit 0
