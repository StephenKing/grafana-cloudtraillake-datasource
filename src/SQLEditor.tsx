import { SQLEditor as SQLCodeEditor } from '@grafana/experimental';
import { DataSource } from 'datasource';
import React, { useRef, useEffect } from 'react';
import {CtlQuery} from 'types';

interface RawEditorProps {
  query: CtlQuery;
  onRunQuery?: () => void;
  onChange: (q: CtlQuery) => void;
  datasource: DataSource;
}

export default function SQLEditor({ query, datasource, onRunQuery, onChange }: RawEditorProps) {
  const queryRef = useRef<CtlQuery>(query);
  useEffect(() => {
    queryRef.current = query;
  }, [query]);


  return (
    <SQLCodeEditor
      query={query.rawSQL}
      onChange={(rawSQL) => onChange({ ...queryRef.current, rawSQL })}
    ></SQLCodeEditor>
  );
}

