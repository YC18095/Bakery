import React from 'react';
import { Badge } from 'react-bootstrap'

function BadgeStatus(props) {

  const { status } = props

  const getBadge = status => {
    switch (status) {
      case 1:
        return <Badge bg='success'>Active</Badge>
      case 2:
        return <Badge bg='secondary'>Unactive</Badge>
      default:
        return <Badge bg='danger'>Unknown</Badge>
    }
  }

  return (
    getBadge(status)
  );
}

export default BadgeStatus;