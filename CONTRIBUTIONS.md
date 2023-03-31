<h3> How to run locally? </h3>

You must have <a href="https://www.docker.com"> Docker </a> installed on your machine.

<p>
    <ul>
        <li>Clone the repository</li>
        <li>Run <code> make gen_and_run </code> </li>
    </ul>
</p>

Voila!

<code> make gen_and_run </code> will run <code>scrap_proto.go</code> inside a docker 
and scrap the proto-related file from <a href="https://github.com/Via-Go/proto">proto repository</a>
and then run <code> docker-compos-dev.yaml </code> <br>
It starts the following services:

- ScyllaDB, on port 9042
- Prometheus, on port 9090
- Grafana, on port 3000
- UsersMS, gRPC server on port 50051 and its REST metrics on port 7070

<a href="https://youtu.be/z5wSmeOoYTI"> Here </a> is short youtube video with presentation of cloning the repository, starting the microservice and exploring
its features.
