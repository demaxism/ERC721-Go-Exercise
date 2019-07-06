const Migrations = artifacts.require("Migrations");
const StarNFT = artifacts.require('./StarNFT');

module.exports = async function(deployer) {
  deployer.deploy(Migrations);
};
