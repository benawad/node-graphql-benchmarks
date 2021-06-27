# TL;DR

- graphql-jit helps
- apollo-server adds overhead
- tracing resolvers kills performance
- type-graphql adds overhead

# Explanation

For further details, please check out [this video](https://www.youtube.com/watch?v=JbV7MCeEPb8).

# Usage

```
git clone https://github.com/benawad/benchmarks
cd benchmarks
npm install
npm start
```

# Benchmarks
duration: 5s
connections: 5
pipelining: 1

| Server                                                                                                                                                                  | Requests/s | Latency | Throughput/Mb |
| :--                                                                                                                                                                     | --:        | :-:     | --:           |
| [go-graphql](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/go-graphql.js)                                                                   | 49168.0    | 0.01    | 81.68         |
| [uWebSockets-graphql+jit](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/uWebSockets-graphql+jit.js)                                         | 13061.6    | 0.02    | 80.36         |
| [core-graphql-jit-str](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/core-graphql-jit-str.js)                                               | 9383.2     | 0.04    | 58.33         |
| [core-graphql-jit-buf](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/core-graphql-jit-buf.js)                                               | 9272.8     | 0.04    | 57.65         |
| [core-graphql-jit-buf-fjs](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/core-graphql-jit-buf-fjs.js)                                       | 9144.8     | 0.04    | 56.83         |
| [benzene-http](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/benzene-http.js)                                                               | 9050.4     | 0.04    | 56.69         |
| [mercurius+graphql-jit](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/mercurius+graphql-jit.js)                                             | 8790.0     | 0.04    | 55.03         |
| [fastify-REST](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/fastify-REST.js)                                                               | 7885.2     | 0.07    | 63.23         |
| [mercurius+graphql-compose](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/mercurius+graphql-compose.js)                                     | 5355.6     | 0.29    | 33.54         |
| [express-REST](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/express-REST.js)                                                               | 5333.2     | 0.16    | 43.10         |
| [graphql-api-koa+graphql-jit](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/graphql-api-koa+graphql-jit.js)                                 | 5306.0     | 0.28    | 33.18         |
| [mercurius](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/mercurius.js)                                                                     | 5249.2     | 0.30    | 32.87         |
| [mercurius+graphql-jit+type-graphql](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/mercurius+graphql-jit+type-graphql.js)                   | 5040.4     | 0.31    | 31.56         |
| [express-gql](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/express-gql.js)                                                                 | 4674.0     | 0.35    | 29.56         |
| [fastify-express-graphql-jit](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/fastify-express-graphql-jit.js)                                 | 4658.8     | 0.36    | 1.28          |
| [fastify-express-graphql-typed-jit](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/fastify-express-graphql-typed-jit.js)                     | 4375.0     | 0.41    | 1.20          |
| [express-graphql+graphql-jit+graphql-compose](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/express-graphql+graphql-jit+graphql-compose.js) | 4058.0     | 0.48    | 25.67         |
| [express-graphql+graphql-jit](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/express-graphql+graphql-jit.js)                                 | 4038.0     | 0.48    | 25.54         |
| [graphql-api-koa](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/graphql-api-koa.js)                                                         | 3790.2     | 0.74    | 23.70         |
| [express-graphql+graphql-jit+type-graphql](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/express-graphql+graphql-jit+type-graphql.js)       | 3727.0     | 0.99    | 23.57         |
| [apollo-server-koa+graphql-jit+type-graphql](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/apollo-server-koa+graphql-jit+type-graphql.js)   | 3674.6     | 0.67    | 23.00         |
| [fastify-express-grapql-typed](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/fastify-express-grapql-typed.js)                               | 3492.6     | 1.25    | 0.96          |
| [express-graphql](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/express-graphql.js)                                                         | 3198.6     | 1.27    | 20.22         |
| [express-graphql+graphql-compose](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/express-graphql+graphql-compose.js)                         | 3138.2     | 1.27    | 19.85         |
| [express-graphql+type-graphql](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/express-graphql+type-graphql.js)                               | 3076.6     | 1.27    | 19.45         |
| [apollo-schema+async](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/apollo-schema+async.js)                                                 | 3068.2     | 1.27    | 19.40         |
| [express-graphql-dd-trace-no-plugin](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/express-graphql-dd-trace-no-plugin.js)                   | 3049.8     | 1.27    | 19.28         |
| [graphql-compose+async](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/graphql-compose+async.js)                                             | 3010.6     | 1.28    | 19.04         |
| [type-graphql+middleware](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/type-graphql+middleware.js)                                         | 2928.6     | 1.27    | 18.52         |
| [type-graphql+async](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/type-graphql+async.js)                                                   | 2777.4     | 1.32    | 17.56         |
| [type-graphql+async-middleware](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/type-graphql+async-middleware.js)                             | 2755.8     | 1.34    | 17.43         |
| [apollo-server-express](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/apollo-server-express.js)                                             | 2486.2     | 1.48    | 15.80         |
| [apollo-opentracing](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/apollo-opentracing.js)                                                   | 1700.0     | 2.70    | 10.81         |
| [express-graphql-dd-trace-less](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/express-graphql-dd-trace-less.js)                             | 1562.0     | 2.64    | 9.88          |
| [express-graphql-dd-trace](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/express-graphql-dd-trace.js)                                       | 1530.6     | 2.81    | 9.68          |
| [apollo-server-express-tracing](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/apollo-server-express-tracing.js)                             | 1505.6     | 2.87    | 47.90         |
| [yoga-graphql-trace](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/yoga-graphql-trace.js)                                                   | 185.0      | 30.14   | 0.08          |
| [apollo-server-fastify+graphql-jit](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/apollo-server-fastify+graphql-jit.js)                     | 0.0        | 0.00    | 0.00          |
| [apollo-server-fastify](https://github.com/benawad/node-graphql-benchmarks/tree/master/benchmarks/apollo-server-fastify.js)                                             | 0.0        | 0.00    | 0.00          |
