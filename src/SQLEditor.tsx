import { SQLEditor as SQLCodeEditor } from '@grafana/experimental';
import { DataSource } from 'datasource';
import React, {useRef, useEffect, useCallback, useMemo} from 'react';
import {CtlQuery} from 'types';
import {getCtlCompletionProvider} from "./language/completionItemProvider";

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

//  const interpolate = (value: string | undefined) => {
//    if (!value) {
//      return value;
//    }
//
//    value = value.replace(EDSID_MACRO, queryRef.current.edsId ?? '');
//    value = getTemplateSrv().replace(value);
//
//    return value;
//  };

  const getEventDataStores = useCallback(async () => {
    const eds: string[] = await datasource.getEventDataStores().catch(() => []);
    return eds.map((eds) => ({ name: eds, completion: eds }));
  }, [datasource]);

  const completionProvider = useMemo(
      () => getCtlCompletionProvider({ getEventDataStores }),
      [getEventDataStores]
  );

  return (
    <SQLCodeEditor
      query={query.rawSQL}
      onChange={(rawSQL) => onChange({ ...queryRef.current, rawSQL })}
      language={{
        id: 'sql',
        completionProvider,
      }}

    ></SQLCodeEditor>
  );
}

