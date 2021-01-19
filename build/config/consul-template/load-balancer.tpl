upstream grpc_servers {
   {{ range $i, $e := service "hello_server|passing" }}
   server {{ $e.Address }}:{{ $e.Port }};
   {{ end }}
}

server {
   listen      5000 http2;
   server_name localhost;

   location /protos.Ping/SayHello {
      grpc_pass grpc://grpc_servers;
   }
}