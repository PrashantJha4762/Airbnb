
module.exports = {
  async up (queryInterface, Sequelize) {
    await queryInterface.sequelize.query(`Create table Idempotency(
      id INTEGER not null primary key auto_increment,
      idemkey varchar(255) not null,
      Created_At datetime not null default current_timestamp,
      Updated_At datetime not null default current_timestamp on update current_timestamp,
      finalized boolean not null default false,
      bookingId integer not null unique,
      foreign key(bookingId) references Bookings(id) )`)
  },

  async down (queryInterface, Sequelize) {
    await queryInterface.sequelize.query('drop table Idempotency')
  }
};
