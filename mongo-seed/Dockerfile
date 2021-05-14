FROM stefanwalther/mongo-seed
ENV MONGODB_HOST mongo
ENV MONGODB_PORT 27017
COPY mongo_sample/people_descriptions.csv ./
CMD ["mongoimport -h mongo -u root -p root --port 27017 --authenticationDatabase admin --db sample --collection people_description --type csv --headerline --file ./people_descriptions.csv"]