module.exports = {
  async up(queryInterface) {
    await queryInterface.sequelize.query(`
      ALTER TABLE hotel
      ADD COLUMN deleted_at DATETIME NULL DEFAULT NULL
    `);
  },

  async down(queryInterface) {
    await queryInterface.sequelize.query(`
      ALTER TABLE hotel
      DROP COLUMN deleted_at
    `);
  }
};