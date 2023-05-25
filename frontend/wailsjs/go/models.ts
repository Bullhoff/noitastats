export namespace main {
	
	export class SessionJsonStruct {
	    favorite: boolean;
	    comment: string;
	
	    static createFrom(source: any = {}) {
	        return new SessionJsonStruct(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.favorite = source["favorite"];
	        this.comment = source["comment"];
	    }
	}

}

