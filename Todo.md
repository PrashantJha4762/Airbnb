Lecture 19:Integrating ORM with Express.js 

Install sequelize mysql2 and sequelize cli from website
make a .sequelizerc file and use common js module.exports to export the paths for config, models, migrations and seeders name them as config, modelpath, migrationpath and seederpath respectively and make sure these files are in the db folder
Eveything that was in the json file should be moved to the config.ts file
in index.ts in the config folder just like you made for serverconfig now made things for db config make sure you have the database 
generate the npx sequelize-cli migration:generate --name create-hotels-table and write the query using sequelize query interface to create the hotels table with the following columns id, name, description, location, price, createdAt and updatedAt
now run the migration using npx sequelize-cli db:migrate and check if the table is created in the database

Lectrure 20: Writing apis with sequelize and express.js