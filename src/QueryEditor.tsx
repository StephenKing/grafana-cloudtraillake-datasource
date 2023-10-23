import React, { useCallback, useEffect, useState } from 'react';
import { QueryEditorProps } from '@grafana/data';
import { QueryEditorHeader } from '@grafana/aws-sdk';
import { CtlDataSourceOptions, CtlQuery } from './types';
import { DataSource } from './datasource';
import { QueryEditorForm } from './QueryEditorForm';

export function QueryEditor(props: QueryEditorProps<DataSource, CtlQuery, CtlDataSourceOptions>) {
    const [dataIsStale, setDataIsStale] = useState(false);
    const { onChange } = props;

    useEffect(() => {
        setDataIsStale(false);
    }, [props.data]);

    const onChangeInternal = useCallback(
        (query: CtlQuery) => {
            setDataIsStale(true);
            onChange(query);
        },
        [onChange]
    );

    return (
        <>
            {props?.app !== 'explore' && (
                <QueryEditorHeader<DataSource, CtlQuery, CtlDataSourceOptions>
                    {...props}
                    enableRunButton={dataIsStale && !!props.query.rawSQL}
                    showAsyncQueryButtons={true}
                    cancel={props.datasource.cancel}
                />
            )}
            <QueryEditorForm {...props} onChange={onChangeInternal} />
        </>
    );
}
