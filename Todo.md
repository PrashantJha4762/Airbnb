Lecture 19:Integrating ORM with Express.js 

Install sequelize mysql2 and sequelize cli from website
make a .sequelizerc file and use common js module.exports to export the paths for config, models, migrations and seeders name them as config, modelpath, migrationpath and seederpath respectively and make sure these files are in the db folder
Eveything that was in the json file should be moved to the config.ts file
in index.ts in the config folder just like you made for serverconfig now made things for db config make sure you have the database 
generate the npx sequelize-cli migration:generate --name create-hotels-table and write the query using sequelize query interface to create the hotels table with the following columns id, name, description, location, price, createdAt and updatedAt
now run the migration using npx sequelize-cli db:migrate and check if the table is created in the database

Lecture 20: Writing apis with sequelize and express.js

Add a coloumn named rating by generating a new migrration
Add a model for your databse in the models folder . use the website for documentation on how to create a model using sequelize cli
after creating the class there is something called as init method which is used to initialize the model with the database connection and define the attributes of the model. Make sure to define the attributes according to the columns in your hotels table.
Then we have to create a sequelize configuration object which would be quite same as the config folder code. The config folder code was basically the configuration written for migrations
After this use sequelize.authenticate() method in server.ts file and create a hotel there using the create method
After this first make a dto layer and then repository layer. Since we will be using the transfer object which will initiate a network request from the clinet , we will make the dto layer.In the repository layer use async functions
Now write the service layer and then controllers and then router and then validation layer
And implement getAllhandler it's a part of hw

Lecture 21:Tombstone in database

Install a library http status codes(search on web) 
create a migration to add a column named Deleted_At in the hotels table and set its default value to null and make sure to set the allowNull property to true(way for adding soft delete functionality in the database )
update the model
Implement the function named Softdelete in the repository layer which will update the Deleted_At column with the current timestamp for a given hotel id (dont forget to save the changes in the database using the save method)
Also modify the function named getAllHotels in the repository layer to return only those hotels which have Deleted_At column as null use where clause in the findAll method of sequelize to filter out the deleted hotels. similarily write the service layer and controlller layer and then write the router layer