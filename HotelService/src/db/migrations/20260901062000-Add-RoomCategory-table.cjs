module.exports = {
  async up (queryInterface) {
    await queryInterface.sequelize.query(`CREATE TABLE room_categories (
      id INT NOT NULL AUTO_INCREMENT,
      hotel_id INT NOT NULL,
      price DECIMAL(10,2) NOT NULL,
      room_type ENUM('single', 'double', 'suite') NOT NULL,
      room_count INT NOT NULL,
      created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
      updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
      deleted_at DATETIME DEFAULT NULL,
      PRIMARY KEY (id),
      FOREIGN KEY (hotel_id) REFERENCES hotel(id)
    )`);
  },

  async down (queryInterface) {
    await queryInterface.sequelize.query(`DROP TABLE room_categories`);
  }
};
