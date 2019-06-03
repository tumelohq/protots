import {TestMessage, WithoutComment, EmptyType} from "./message";
import {ProtoAPIService} from "@tumelohq/protots-base"

// TestService comments for testing
export interface TestService {

  // TestEndpointGet comment for testing
  TestEndpointGet(arg: TestEndpointRequest): Promise<TestEndpointResponse>

  TestEndpointPost(arg: TestEndpointRequest): Promise<TestEndpointResponse>

}


// TestService comments for testing
export class TestService extends ProtoAPIService implements TestService {

  //  TestEndpointGet comment for testing
  TestEndpointGet(arg: TestEndpointRequest): Promise<TestEndpointResponse> {
    const u = "/testbefore/" + arg.id + "/after"
    return this.get(u, arg)
  }


  TestEndpointPost(arg: TestEndpointRequest): Promise<TestEndpointResponse> {
    const u = "/testbefore/" + arg.id + "/after"
    return this.post(u, arg)
  }

}


export interface TestEndpointRequest {
  id: string,
}


export interface TestEndpointResponse {
  message: TestMessage,
}
