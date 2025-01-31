# Unit Testing a Go net/http API built with Entgo.io

## Introduction

Unit testing APIs, especially those that involve CRUD operations, has often required multiple third party testing libraries and frameworks due to language-specific limitations, inherently being part of client/server interactions, and the system integration mocking involved.  Often, teams need a hybrid setup that rely on services like a database or webserver, or just fall back to integration testing.

With Go and some sensible design, true unit testing becomes a lot easier.  New features added to Go's net/http package in 1.22, we no longer require third party libraries like Gin or Gorilla in our application or our testing strategy.  This sample project demonstrates how unit testing can work on a vanilla Go REST application.  We use Ent as an ORM, but the principles can be applied to ORM-less applications.