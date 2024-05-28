import React, { useState, useEffect } from 'react';
import Box from '../box';
import { Ratio, Image } from 'react-bootstrap';
import { Link } from 'react-router-dom';
import BadgeStatus from '../public/badge-status';

function Relation(props) {

  const { typeId } = props

  const [ data, setData ] = useState(null)

  const [ page, setPage ] = useState(0)

  useEffect(() => {
    fetch('http://localhost:8000/api/admin/product/listbytype/' + typeId + '/5/' + page)
      .then(response => {
        if (!response.ok) {
          console.log('error')
        }
        return response.json()
      })
      .then(json => {
        setData(json.data)
      })
  }, [page])

  return (
    <Box.Aside title='Related Items' type='Products' page={page} setPage={setPage} total={data ? Math.ceil(data.total / 5) : 0}>
      {data ? data.list.map((item, index) => {
        return (
          <div key={index} className='item py-2'>
            <Ratio aspectRatio='1x1'>
              <Image src={window.location.origin + `/images/` + item.images[0]} alt={item.name} />
            </Ratio>
            <div className='content ps-2'>
              <h3 className='mb-0'>
                <Link to={`/admin/product/detail/${item.id}`} className='stretched-link'>
                  <span className='d-block text-truncate'>{item.name}</span>
                </Link>
              </h3>
              <div className='status text-end'>
                <BadgeStatus status={item.status} />
              </div>
            </div>
          </div>
        );
      })
      : ''}
    </Box.Aside>
  );
}

export default Relation;