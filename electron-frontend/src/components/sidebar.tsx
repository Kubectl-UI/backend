import React, { useState } from 'react';

export default function Sidebar() {
  const [list] = useState(['Pods', 'Deployments', 'Services', 'namespaces']);
  return <div> nav </div>;
}
