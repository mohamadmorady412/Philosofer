# Philosofer

![Kubernetes](https://img.shields.io/badge/Kubernetes-v1.20-blue?logo=kubernetes)

The system's populace is tripartite: users, data procurers, and administrators, each endowed with the capacity for ingress, registration, and egress.**

Users possess the faculty to inscribe "posts" – articulations of predicaments encountered by themselves or their organizations in the realm of commerce. Furthermore, they may append "approaches" – elucidations of resolutions – to the posts of fellow users or generate them autonomously. They are also empowered to confer "donations" upon the solutions of others or orchestrate "donation" campaigns for their own approaches. Users may convene "events," gatherings intended for the execution of tasks (e.g., the advancement of proposals aimed at value creation), which may also feature donation drives.

Users can extend or receive "job offers" to or from other users. They may upload "resumes," compose "descriptions," and engage in other customary activities, such as the uploading of images, within their accounts.

All user-generated data, encompassing posts, approaches, and other pertinent information, is abstracted and transferred to a "data lake," ensuring the preservation of user privacy.

Data procurers, upon remittance of requisite fees, are granted access to the data lake's API.

Users may establish "follow" relationships with other accounts, with lists of "followers" and "following" displayed within the "follow" section, enabling awareness of their activities through notifications.

Users may engage in "exploration," wherein they are presented with "approaches," "events," and "posts" from other users, curated by a "suggestion" algorithm that considers their exploration history, resumes, followed users, and other relevant factors.

## Key Features

* **Users:**
    * The system's populace is tripartite: users, data procurers, and administrators, each endowed with the capacity for ingress, registration, and egress.
* **User Capabilities:**
    * Users possess the faculty to inscribe "posts" – articulations of predicaments encountered by themselves or their organizations in the realm of commerce.
    * Furthermore, they may append "approaches" – elucidations of resolutions – to the posts of fellow users or generate them autonomously.
    * They are also empowered to confer "donations" upon the solutions of others or orchestrate "donation" campaigns for their own approaches.
    * Users may convene "events," gatherings intended for the execution of tasks (e.g., the advancement of proposals aimed at value creation), which may also feature donation drives.
    * Users can extend or receive "job offers" to or from other users.
    * They may upload "resumes," compose "descriptions," and engage in other customary activities, such as the uploading of images, within their accounts.
* **Data Lake:**
    * All user-generated data, encompassing posts, approaches, and other pertinent information, is abstracted and transferred to a "data lake," ensuring the preservation of user privacy.
* **Data Access:**
    * Data procurers, upon remittance of requisite fees, are granted access to the data lake's API.
* **Follow:**
    * Users may establish "follow" relationships with other accounts, with lists of "followers" and "following" displayed within the "follow" section, enabling awareness of their activities through notifications.
* **Explore:**
    * Users may engage in "exploration," wherein they are presented with "approaches," "events," and "posts" from other users, curated by a "suggestion" algorithm that considers their exploration history, resumes, followed users, and other relevant factors.

```bash
project-root/
├── services/
│   ├── auth-service/ (golang)
│   ├── user-service/ (golang)
│   ├── post-service/ (golang)
│   ├── event-service/ (golang)
│   ├── job-service/ (golang)
│   ├── notification-service/ (golang)
│   ├── explore-service/ (golang + ElasticSearch/Pinecone)
│   ├── data-lake-service/ (golang + Cassandra)
│   ├── gateway/ (API Gateway - Nginx/Kong)
│   ├── common/ (Shared packages)
│   ├── Dockerfile (Each service separately)
├── data-pipeline/ (data pricces)
│   ├── stream-processing/ (Kafka Streams + Flink/Golang)
│   ├── batch-processing/ (Spark + Python)
│   ├── recommendation/ (golang + TensorFlow/Pinecone)
│   ├── Dockerfile
├── database/
│   ├── migrations/
│   ├── init-scripts/
│   ├── managed-db/
│   ├── nosql-db/ (Cassandra for no structured data)
├── infrastructure/
│   ├── kubernetes/
│   ├── docker-compose.yml
│   ├── monitoring/ (Monitoring and logs)
│   │   ├── prometheus/
│   │   ├── grafana/
│   │   ├── alertmanager/
│   ├── failover/ (Failover management for databases and services)
│   ├── cdn/
│   ├── serverless/
├── tests/
│   ├── integration/
│   ├── e2e/
│   ├── load-testing/ (Scalability test with K6 and Locust)
├── docs/
├── .gitignore
├── README.md
```