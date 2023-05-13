import vest, { test, enforce } from 'vest'
import { isEmail } from 'validator'

// https://codesandbox.io/s/svelte-vest-form-validation-example-kf185?file=/validate.js:0-1387
// extend vest with customer email validator
enforce.extend({ isEmail })

const suite = vest.create('user_form', (data = {}, currentField) => {
  // if field name is supply validate only that field
  vest.only(currentField)

  test('username', 'Username is required', () => {
    enforce(data.username).isNotEmpty()
  })

  test('username', 'Username must be at least 3 characters long', () => {
    enforce(data.username).longerThanOrEquals(3)
  })

  test('email', 'Email address is not valid', () => {
    enforce(data.email).isEmail()
  })

  test('password', 'Password is weak, Maybe add a number?', () => {
    // mark this test as a warning only
    vest.warn()
    enforce(data.password).matches(/[0-9]/)
  })

  if (data.password) {
    test('confirm_password', 'Passwords do not match', () => {
      enforce(data.confirm_password).equals(data.password)
    })
  }

  test('tos', () => {
    enforce(data.tos).isTruthy()
  })
})

export default suite
