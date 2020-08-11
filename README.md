# Workshop Description: Infrastructure for a Fraud Detection ML Application

This is a hands on workshop that given at the [Go Get conference](https://www.gogetcommunity.com/) in 05/2020 and will be given at [GopherCon](https://www.gophercon.com/) in 11/2020.

The architecture is based on an architecture used in production by PayPal for fraud detection utilizng an ML model. 
The workshop is designed for the conference attendees to have a hands on experience of a ML infrastructure. As it based on real-life credit card transactions, there are abstractions in place, especially around reading/writing the transaction content, to allow maximal utilization of the data.

## Setup

- Have [GitHub](https://help.github.com/en/github/getting-started-with-github/set-up-git) set up on your machine
- Install [Go](https://golang.org/dl/) (v1.14)
- Install latest [Docker](https://docs.docker.com/get-docker/)
- Pull latest Docker images:
  - docker pull aerospike/aerospike-server
  - docker pull aerospike/aerospike-tools
  - docker pull tensorflow/serving
- Download the [Kaggle dataset](https://www.kaggle.com/mlg-ulb/creditcardfraud/data) of credit card transactions 


In case you are new to Go, here are some useful links:
- A summarized overview of Go: [Learn Go in Y Minutes](https://learnxinyminutes.com/docs/go/)
- A short hands on tutorial: [Tour of Go](http://tour.golang.org/) 
- A handy cheatsheet: [Go by Example](https://gobyexample.com/)
