import {GRPCCodes} from "./codes";
import ExtendableError from 'ts-error';

/**
 *  GRPC Status is the content of the error message from the grpc service
 */
export interface GRPCStatus {
  error: string
  message: string
  code: GRPCCodes
  details?: any[]
}

/**
 *
 */
export class GRPCError extends ExtendableError {
  readonly grpcStatus: GRPCStatus

  constructor(errorBody: GRPCStatus) {
    super(`${JSON.stringify(errorBody)}`)
    this.grpcStatus = errorBody
  }
}


/**
 *
 */
export class HTTPError extends ExtendableError {

}