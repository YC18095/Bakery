import React, { useState, useEffect } from 'react';
import Search from '../public/search';
import { Form } from 'react-bootstrap';

function ProductFeature(props) {

  const { keyword, setKeyword, typeId, setTypeId, eventIds, setEventIds } = props

  const [typeList, setTypeList] = useState(null)

  const [eventList, setEventList] = useState(null)

  useEffect(() => {
    Promise.all([
      fetch(`${process.env.REACT_APP_API_URL}/type/list/0/0`),
      fetch(`${process.env.REACT_APP_API_URL}/event/list/0`)
    ])
      .then(([typeResponse, eventResponse]) => {
        if (!typeResponse.ok && !eventResponse.ok) {
          console.log('error')
        }
        typeResponse.json().then((json) => setTypeList([{ id: 0, name: 'All' }, ...json.data.list]))
        eventResponse.json().then((json) => setEventList(json.data.list))
      })
  }, [])

  const handleChangeEvent = eventId => {
    let active = false
    let eventMap = []
    for (let i = 0; i < eventIds.length; i++) {
      if (eventId >= eventIds[i] && !active) {
        active = true
        if (eventId > eventIds[i]) eventMap.push(eventId)
        else if (eventId === eventIds[i]) i++
      }
      eventMap.push(eventIds[i])
    }
    setEventIds(eventMap)
  }

  return (
    <Form className='product__filter'>
      <Search keyword={keyword} setKeyword={setKeyword} className='mb-3' />
      {typeList ? 
      <div className='mb-3'>
        <h6 className='mb-2'>Types:</h6>
        {typeList.map((item, index) => {
          return (
            <Form.Check defaultChecked={0} type='radio' label={item.name} key={index} id={`typeInput${item.id}`} name='type' value={item.id} checked={item.id === typeId} onChange={() => setTypeId(item.id)} />
          );
        })}
      </div>
      : ''}
      {eventList ? 
        <div className='mb-3'>
          <h6 className='mb-2'>Event:</h6>
          {eventList && eventList.map((item, index) => {
            return (
              <Form.Check type='checkbox' label={item.name} key={index} id={`eventInput${item.id}`} name='event' value={item.id} checked={eventIds.includes(item.id)} onChange={() => handleChangeEvent(item.id)} />
            );
          })}
        </div>
      : ''}
    </Form>
  );
}

export default ProductFeature;
