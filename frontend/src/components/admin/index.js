import React from 'react';
import { Container, Row, Col } from 'react-bootstrap'
import { Routes, Route } from 'react-router-dom'
import Sidebar from "./sidebar";
import Topbar from "./topbar";
import TypeList from "./type/list";
import TypeDetail from "./type/detail";
import TypeEdit from "./type/edit";

function AdminLayout() {

  return (
    <Container fluid className='px-4 admin'>
      <Row className='g-4'>
        <Sidebar />
        <Col className='py-3'>
          <Topbar />
          <Routes>
            <Route path='type'>
              <Route path='' element={<TypeList />} />
              <Route path='detail/:id' element={<TypeDetail />} />
              <Route path='edit/:id' element={<TypeEdit />} />
            </Route>
          </Routes>
        </Col>
      </Row>
    </Container>
  );
}

export default AdminLayout;