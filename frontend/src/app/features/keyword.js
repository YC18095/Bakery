import { createSlice } from '@reduxjs/toolkit'

const keyword = createSlice({
  name: 'keyword',
  initialState: {
    value: "abc",
  },
  reducers: {
    setValue: (state, value) => {
      state.value = value
    }
  },
})

export const { setValue } = keyword.actions
export const getValue = (state) => state.keyword.value

export default keyword.reducer