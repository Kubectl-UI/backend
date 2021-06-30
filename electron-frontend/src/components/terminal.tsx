import React, { useState } from 'react';

export default function Terminal() {
  const [title] = useState('user@host:/$ _ ');
  return (
    <div className="terminal">
      <div>
        <div className="terminal-controls">
          <div className="buttons">
            <button type="button" className="close">
              x
            </button>
            <button type="button" className="minimize">
              -
            </button>
            <button type="button" className="enlarge">
              +
            </button>
          </div>
          <div className="terminal-user-info">users@host:/$</div>
        </div>
      </div>
      <div className="terminal-body">
        <p>
          jQuery Terminal Emulator v. 2.26.0 (c) 2011-2021 Jakub T. Jankiewicz{' '}
        </p>
        <textarea name="" value={title} id="" cols={30} rows={10} />
      </div>
    </div>
  );
}
