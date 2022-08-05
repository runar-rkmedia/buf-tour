import {
  createConnectTransport,
  createPromiseClient,
  createCallbackClient,
} from "@bufbuild/connect-web";

// Import service definition that you want to connect to.
import { PetStoreService } from "../../gen/proto/web/pet/v1/pet_connectweb";

// The transport defines what type of endpoint we're hitting.
// In our example we'll be communicating with a Connect endpoint.
const transport = createConnectTransport({
  baseUrl: "http://localhost:8080",
});

// Here we make the client itself, combining the service
// definition with the transport.
const client = createPromiseClient(PetStoreService, transport);
const cbClient = createCallbackClient(PetStoreService, transport);

async function testClientType() {
  // looks like this response is not typed
  const response = await client.putPet({ petId: "bob" });

  cbClient.putPet({}, (err, res) => {
    // but this one is???
    res.pet?.name;
  });

  return response;
}

testClientType();
