
# What is this?

A small go command line application which displays statistics 
for analyzing api flexibility and support decisions on major version updates and consolidation of a component library.

## What it does

Counts string of web-component tag names in multiple projects.
It has json config file.
Attempts to parse component-library versions from package.json.
The output is a markdown table with the count of string occurances by projects.
Component tag names are parsed from filenames of a dist folder of the library.

# An example of output:

# Components
| Component | Test project1 v0.0.1 | My second project v0.0.2 |
| --- | --- | --- |
| my-component | 4 | 4 |
* Report generated: 2020-08-28 08:24:48.944914209 +0200 CEST m=+0.001882842

