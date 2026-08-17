module.exports = {
  async up (queryInterface) {
      queryInterface.sequelize.query(`Alter table hotel add column rating decimal(3,2) default 0`)
  },

  async down (queryInterface) {
    queryInterface.sequelize.query(`Alter table hotel drop column rating`)
  }
};
