
module.exports = {
  async up (queryInterface) {
    await queryInterface.sequelize.query(`CREATE TABLE rooms (
      id INT NOT NULL AUTO_INCREMENT,
      room_category_id INT NOT NULL,
      hotel_id INT NOT NULL,
      room_no VARCHAR(255) NOT NULL,
      DateOFAVAIlibility DATE NOT NULL,
      booking_id INT,
      created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
      updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
      deleted_at DATETIME DEFAULT NULL,
      PRIMARY KEY (id),
      FOREIGN KEY (room_category_id) REFERENCES room_categories(id),
      FOREIGN KEY (hotel_id) REFERENCES hotel(id)
    )`);
  },

  async down (queryInterface) {
    await queryInterface.sequelize.query(`DROP TABLE rooms`);
  }
};
