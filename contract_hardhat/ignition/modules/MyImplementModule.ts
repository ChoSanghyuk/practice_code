import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

const myProxyV1Module = buildModule("MyProxyV1Module", (m) => {
    const proxyAdminOwner = m.getAccount(0);
  
    const impl = m.contract("MyImplementV1");

    const proxy = m.contract("MyProxy", [
      impl,
      "0x",
    ]);
  
    // 이건 필요 없는듯
    // const proxyAdminAddress = m.readEventArgument(
    //   proxy,
    //   "AdminChanged",
    //   "newAdmin"
    // );
    // const proxyAdmin = m.contractAt("ProxyAdmin", proxyAdminAddress);
  
    return { proxy };
  });

  const myUpgradeableModule = buildModule("MyUpgradeableModule", (m) => {
    const { proxy } = m.useModule(myProxyV1Module);

    /*
    Finally, we'll use the m.contractAt(...) method to tell Ignition to use the ProxyAdmin ABI for the contract at the address we just retrieved. 
    This will allow us to interact with the ProxyAdmin contract when we upgrade our proxy.
    */
    const impl = m.contractAt("MyImplementV1", proxy);

    return { impl, proxy };
});

export default myUpgradeableModule;