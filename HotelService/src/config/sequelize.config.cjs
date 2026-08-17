require('ts-node/register'); // This line enables TypeScript support for Sequelize CLI
const config=require('./db.config.cjs')
module.exports=config;
