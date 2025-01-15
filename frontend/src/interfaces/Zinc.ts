export interface ZincResponse {
    took: number;
    timed_out: boolean;
    _shards: Shards;
    hits: Hits;
}

export interface Shards {
    total: number;
    successful: number;
    skipped: number;
    failed: number;
}

export interface Hits {
    total: Total;
    max_score: number;
    hits: Hit[];
}

export interface Total {
    value: number;
}

export interface Hit {
    _index: string;
    _type: string;
    _id: string;
    _score: number;
    "@timestamp": string;
    _source: EmailSource;
}

export interface EmailSource {
    "@timestamp": string;
    Body: string;
    CC: string[];
    "Content-Type": string;
    Date: number;
    From: string;
    "Message-Id": string;
    "Mime-Version": string;
    Subject: string;
    To: string[];
    "X-Filename": string;
    "X-Folder": string;
    "X-From": string;
    "X-Origin": string;
    "X-To": string[];
    "X-bcc": string[];
    "X-cc": string[];
}