module.exports = {
  async up (queryInterface) {
    queryInterface.sequelize.query(`Create table if not exists hotels(
      id int primary key auto_increment,
      name varchar(255) not null,
      address text not null,
      location varchar(255) not null,
      createdAt datetime not null default current_timestamp,
      updatedAt datetime not null default current_timestamp on update current_timestamp
    )`);
  },

  async down (queryInterface) {
    queryInterface.sequelize.query(`drop table if exists hotels`);
  }
};
