# Outputs

Clingy needs an output directory (by default it will use `./output`) to execute. This then will 
be used to store artifacts it generates. If passed a `-d` (debug flag), logs in a `logs.txt` will 
also be stored in that location.

This document reviews various output formats clingy will dump to the `./output/<TIMESTAMP>` path.

The examples generated in this file correspond to this 
[clingy YAML file](https://github.com/madhuravius/clingy/blob/main/cmd/test_data/01_basic_flow_will_pass.yaml).

## HTML Reports - Simple

Example can be found at this link [here](/clingy/example-outputs/html-simple).

This generates a simple HTML report that includes:

* Label
* Description
* For each step (presented as content to scroll through):
  * Label
  * Image
  * Description

## HTML Reports - Carousel

Example can be found at this link [here](/clingy/example-outputs/carousel).

This generates a simple HTML report that includes:

* Label
* Description
* For each step (presented in a slideshow):
    * Label
    * Image
    * Description

## Images Only


This generates a series of images with the following baked into the image:

* Label
* Description

An Example is shown below.

<figure markdown>
  ![Image title](/clingy/example-outputs/images-only/3.jpg)
  <figcaption>Generated image from this report</figcaption>
</figure>