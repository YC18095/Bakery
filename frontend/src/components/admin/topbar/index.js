import React from 'react';
import { Dropdown, Ratio, Image } from 'react-bootstrap'

function Topbar() {
  return (
    <div className='topbar mb-5'>
      <h6 className='mb-0'>Title</h6>
      <div className='topbar__option'>
        <Dropdown align='end'>
          <Dropdown.Toggle variant='none'>
            <span className='text-truncate d-inline-block text-end pe-2'>Jason Cheung</span>
            <Ratio aspectRatio='1x1'>
              <Image src={require(`../../../images/user.png`)} />
            </Ratio>
          </Dropdown.Toggle>
          <Dropdown.Menu className='py-0'>
            <Dropdown.Item className='text-end' href='#'>Sign out<i className="bi bi-box-arrow-left ps-2"></i></Dropdown.Item>
          </Dropdown.Menu>
        </Dropdown>
      </div>
    </div>
  );
}

export default Topbar;