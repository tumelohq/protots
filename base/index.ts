import {DateTimeFormatter, LocalDate, ZonedDateTime} from 'js-joda'
import {GRPCError, GRPCStatus, HTTPError} from "./error";

export type RFC3339String = string

enum HTTPMethod {
  Get = 'GET',
  Post = 'POST',
}

// TODO Add ability to get headers
export class ProtoAPIService {
  private readonly baseURL: string

  constructor(baseURL: string) {
    this.baseURL = baseURL
  }

  get = async <A, T>(endpoint: string, args: A): Promise<T> => {
    const endpointWithParams = ProtoAPIService.substituteEndpointParams(args, endpoint)
    const request = this.createRequest(HTTPMethod.Get, endpointWithParams)
    return this.performRequest<T>(request, endpointWithParams)
  }

  post = async <B, T>(endpoint: string, body: B): Promise<T> => {
    const request = this.createRequest(HTTPMethod.Post, endpoint, body)
    return this.performRequest<T>(request, endpoint)
  }

  createRequest = async <B>(method: HTTPMethod, endpoint: string, body?: B): Promise<Response> => {
    return fetch(this.baseURL + endpoint, {
      method,
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
      },
      body: body === undefined ? undefined : JSON.stringify(body),
    })
  }

  async performRequest<T>(request: Promise<Response>, endpoint: string): Promise<T> {
    try {
      const response = await request
      if (response.status == 200) {
        const content = response.json()
        return content as Promise<T>
      } else {
        return this.handleError(response)
      }
    } catch (e) {
      throw new Error(`Error making request ${this.baseURL}${endpoint}: ${e.message}`)
    }
  }

  handleError = async (response: Response) => {
    let errorToThrow: HTTPError | GRPCError
    try {
      const errorBody: GRPCStatus = await response.json()
      errorToThrow = new GRPCError(errorBody)
    } catch {
      let body: string | undefined = response.body === null ? undefined : response.body.toString()
      errorToThrow = new HTTPError(response.status, body)
    }
    throw errorToThrow
  }

  static RFC3339FromLocalDate = (localDate: LocalDate): RFC3339String => {
    return DateTimeFormatter.ISO_LOCAL_DATE.format(localDate) + 'T00:00:00Z'
  }

  static ZonedDateTimeFromRFC3339 = (zonedTime: RFC3339String,): ZonedDateTime => {
    return ZonedDateTime.parse(zonedTime)
  }

  static LocalDateFromRFC3339 = (zonedTime: RFC3339String): LocalDate => {
    const time = ZonedDateTime.parse(zonedTime)
    return LocalDate.from(time)
  }

  // If the args contain a property name that exists in the endpoint surrounded by { },
  // the value of the property in the endpoint will be replace, otherwise the param and
  // its value is added as a URL parameter. Returns a new endpoint with args substituted
  // with the URL params appended
  // TODO NEED TO CHECK THAT ENDPOINTS ARE NOT ALREADY IN THE URL
  public static substituteEndpointParams = <A>(args: A, endpoint: string): string => {
    let urlParams = ''
    let endpointResult = endpoint
    for (const arg in args) {
      const argVariableName = `{${arg}}`
      const argValue = args[arg]
      if (argValue !== undefined) {
        const argString = `${argValue}`
        if (endpointResult.includes(argVariableName)) {
          endpointResult = endpointResult.replace(argVariableName, argString)
        } else {
          const separator = urlParams.length == 0 ? '' : '&'
          urlParams += `${separator}${arg}=${argString}`
        }
      }
    }
    return urlParams.length > 0 ? `${endpointResult}?${urlParams}` : endpointResult
  }
}

/**
 * Converts JSON-GRPC boolean values that if false are not sent and converts them to a proper boolean
 *
 * @param value
 */
export const HandleBooleanValue = (value: boolean | undefined): boolean => {
  return value === undefined ? false : value
}
