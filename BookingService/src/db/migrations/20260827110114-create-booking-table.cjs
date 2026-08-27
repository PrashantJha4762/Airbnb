module.exports = {
  async up (queryInterface) {
    await queryInterface.sequelize.query(`Create table Bookings(
      id INTEGER PRIMARY KEY AUTO_INCREMENT,
      userId INTEGER NOT NULL,
      hotelId INTEGER NOT NULL,
      CreatedAt DATETIME NOT NULL default current_timestamp,
      UpdatedAt DATETIME NOT NULL default current_timestamp on update current_timestamp,
      bookingAmount integer NOT NULL,
      status ENUM('PENDING', 'CONFIRMED', 'CANCELLED') NOT NULL default 'PENDING',
      totalguest integer NOT NULL
    )`);
  },

  async down (queryInterface, Sequelize) {
    await queryInterface.sequelize.query('DROP TABLE Bookings');
  }
};
