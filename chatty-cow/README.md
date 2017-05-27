chatty-cow
========

Small program design to be ran inside a container in Rancher that exposes an HTTP to be used as a healthcheck. Reports healthy if it can find another chatty-cow container via Rancher metadata and make successful HTTP requests to it.

## Building

`make`


## Running

`./bin/chatty-cow`

## License
Copyright (c) 2014-2016 [Rancher Labs, Inc.](http://rancher.com)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

[http://www.apache.org/licenses/LICENSE-2.0](http://www.apache.org/licenses/LICENSE-2.0)

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
