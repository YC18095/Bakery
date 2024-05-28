import { configureStore } from '@reduxjs/toolkit'
import keyword from './features/keyword'

export default configureStore({
  reducer: {
    keyword: keyword
  }
})