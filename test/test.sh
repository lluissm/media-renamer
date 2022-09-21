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

validate_initial_state() {
    # test files exist
    if [[ ! -f "$PROJECT_DIR/.hidden1.jpeg" ]]; then exit 1; fi
    if [[ ! -f "$PROJECT_DIR/1.jpeg" ]]; then exit 1; fi
    if [[ ! -f "$PROJECT_DIR/2.jpeg" ]]; then exit 1; fi
    if [[ ! -f "$PROJECT_DIR/3.mov" ]]; then exit 1; fi
    if [[ ! -f "$PROJECT_DIR/4.mov" ]]; then exit 1; fi
}

validate_renaming() {
    # .hidden1.jpeg is not modified
    if [[ ! -f "$PROJECT_DIR/.hidden1.jpeg" ]]; then exit 1; fi

    # 1.jpeg is renamed properly
    if [[ -f "$PROJECT_DIR/1.jpeg" ]]; then exit 1; fi
    if [[ ! -f "$PROJECT_DIR/2017_01_15_14_10_35.jpeg" ]]; then exit 1; fi
    # 2.jpeg is renamed properly
    if [[ -f "$PROJECT_DIR/2.jpeg" ]]; then exit 1; fi
    if [[ ! -f "$PROJECT_DIR/2021_05_23_08_05_12.jpeg" ]]; then exit 1; fi
    # 3.mov is renamed properly
    if [[ -f "$PROJECT_DIR/3.mov" ]]; then exit 1; fi
    if [[ ! -f "$PROJECT_DIR/2013_12_11_11_12_13.mov" ]]; then exit 1; fi
    # 4.mov is renamed properly
    if [[ -f "$PROJECT_DIR/4.mov" ]]; then exit 1; fi
    if [[ ! -f "$PROJECT_DIR/1999_08_07_19_25_47.mov" ]]; then exit 1; fi
}

run_version_test() {
    output=$($CMD -version)
    if [[ ! "$output" =~ "version:" ]]; then exit 1; fi
}

run_test() {
    path=$1
    delete_sample_project
    extract_sample_project
    validate_initial_state
    $($CMD -v -c "$1" $PROJECT_DIR)
}

# Good custom config file should rename all file
run_test "sample-config.yml"
validate_renaming

# Empty custom config file should not alter any file
run_test "empty-config.yml"
validate_initial_state

# If custom config file is not provided it should use the default one
run_test ""
validate_renaming

# Test -version flag
run_version_test

delete_sample_project
exit 0
