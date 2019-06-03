import {GRPCCodes} from "./codes";
import ExtendableError from 'ts-error';

/**
 *  GRPC Status is the content of the error message from the grpc service
 */
export interface GRPCStatus {
  message: string
  code: GRPCCodes
  details?: any[]
}

/**
 * GRPCError is the error type that is returned in the case of a grpc error
 */
export class GRPCError extends ExtendableError {
  readonly grpcStatus: GRPCStatus

  constructor(errorBody: GRPCStatus) {
    super(`${JSON.stringify(errorBody)}`)
    this.grpcStatus = errorBody
  }
}

/**
 * HTTPError is the error that is returned when there is an error and the error cannot be types into a GRPCError
 */
export class HTTPError extends ExtendableError {
  readonly httpStatusCode: number
  readonly httpBody: string

  constructor(httpStatusCode: number, errorBody: string | undefined) {
    errorBody = errorBody === undefined ? "" : errorBody
    super(`(HTTP ${httpStatusCode}): ${errorBody}`)
    this.httpStatusCode = httpStatusCode
    this.httpBody = errorBody
  }
}
