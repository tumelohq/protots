import {ProtoAPIService} from "@tumelohq/protots-base"

// TestMessage is the test message type
export interface TestMessage {
  id: string,
  boolean: boolean,
  int32type: number,
  int64type: string,
  uint32type: number,
  uint64type: string,
  enum: TestEnum,
}


export interface WithoutComment {
  id: string,
}


export interface EmptyType {
}


// Enum type
export enum TestEnum {
  TEST_ENUM_1 = "TEST_ENUM_1",
  TEST_ENUM_2 = "TEST_ENUM_2",
}
