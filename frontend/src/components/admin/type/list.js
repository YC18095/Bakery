import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import Box from '../box';
import BadgeStatus from '../public/badge-status';

function TypeList() {
  
  const [ typeList, setTypeList ] = useState(null)

  const fields = [
    { name: 'Name' }, 
    { name: 'Status', className: 'text-center status' }, 
    { name: 'Products', className: 'text-center relation' }, 
    { className: 'text-center option' }
  ]

  useEffect(() => {
    fetch('http://localhost:8000/api/admin/type/list/10/0').then(response => {
      if (!response.ok) {
        console.log('error')
      }
      return response.json()
    }).then(json => {
      setTypeList(json.data)
    })
  }, [])

  return (
    <Box>
      <Box.Title title='Type Table' create={`create`} />
      <Box.List fields={fields}>
        {typeList && typeList.list.map((item, index) => {
          return (
            <tr key={index}>
              <td className='name'>
                <strong>{item.name}</strong>
                <span className='d-block'>Update at: {new Date(item.updatedAt).toLocaleDateString('en', { year: 'numeric', month: 'short', day: 'numeric' })}</span>
              </td>
              <td className='text-center status'>
                <BadgeStatus status={item.status} />
              </td>
              <td className='text-center relation'>
                <a href='javascript:;'>{item.productCount} Items</a>
              </td>
              <td className='text-center option'>
                <Link to={`detail/${item.id}`}><i className='bi bi-eye-fill me-1'></i>View</Link>
                <Link to={`edit/${item.id}`}><i className='bi bi-pen-fill me-1'></i>Edit</Link>
              </td>
            </tr>
          );
        })}
      </Box.List>
    </Box>
  );
}

export default TypeList;
